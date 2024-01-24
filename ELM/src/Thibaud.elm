module Thibaud exposing (..)

import Browser
import Array exposing (fromList, get)
import Html exposing (..)
import Html.Attributes exposing (style)
import Http 
import Random 
import Json.Decode exposing (Decoder, string, list, field, map2, map)



main : Program () Model Msg
main =
    Browser.element
        { init = init
        , update = update
        , subscriptions = subscriptions
        , view = view
        }

-- Model
type alias Model = 
    { allWord : List String
    , randomword : Maybe String
    , state : State
    , select : String
    }

type State 
    = Failure
    | Loading
    | Success (List Word)


type alias Word =
    { word : String
    , meanings : List Meaning
    }

-- Meaning alias
type alias Meaning =
    { partOfSpeech : String
    , definitions : List Definition
    }

-- Definition alias
type alias Definition =
    { definition : String
    }

-- Msg
type Msg
    = WordFetched (Result Http.Error String)
    | RandomNumber Int
    | DefinitionsFetched (Result Http.Error (List Word))

-- Init function
init : () -> (Model, Cmd Msg)
init _ =
    ( { allWord = []
      , randomword = Nothing
      , select = Maybe.withDefault "" Nothing 
      , state = Loading
      }
    , Http.get { url = "/thousand_words_things_explainer.txt", expect = Http.expectString WordFetched }
    )


-- Update function
update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        WordFetched (Ok wordText) ->
            let
                wordList =
                    textToList wordText

                index = 
                    Random.generate RandomNumber (Random.int 0 (List.length wordList - 1 ))
            in
                ({ model | allWord = wordList}, index)

        WordFetched (Err _) ->
            ({model | state = Failure}, Cmd.none)

        RandomNumber ind ->
            let 
                choix = get ind (fromList model.allWord)
            in 
            case choix of
                Just mot ->
                    ( {model | randomword = choix}, getDefinitionsCmd mot )

                Nothing ->
                    (model, Cmd.none)

        DefinitionsFetched definitionsResult ->
            case definitionsResult of
                Ok definitions ->
                    ({model | state = Success definitions}, Cmd.none)
                    
                Err _ ->
                    ({ model | state = Failure }, Cmd.none)


-- Subscriptions    
subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none

-- View
view : Model -> Html Msg
view model =
    case model.state of
        Failure ->
            text "Impossible de charger votre livre."

        Loading ->
            text "Chargement..."

        Success fullWords ->
            div []
            [div [] (List.map (\fullWord -> text fullWord.word) fullWords)
            , div [] (List.map viewMeaning fullWords)
            ]

-- Fonction auxiliaire pour afficher une signification (meaning)
viewMeaning : Word -> Html Msg
viewMeaning fullWord =
    div []
        (List.map viewMeaningItem fullWord.meanings)

viewMeaningItem : Meaning -> Html Msg
viewMeaningItem meaning =
    div []
        [ div [] [text (meaning.partOfSpeech)]
        , div [] (List.map viewDefinition meaning.definitions)
        ]

viewDefinition : Definition -> Html Msg
viewDefinition definition =
    div []
        [ div [] [text (definition.definition)]
        ]

-- GetDefinitionsCmd
getDefinitionsCmd : String -> Cmd Msg
getDefinitionsCmd word =
    Http.get
        { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ word
        , expect = Http.expectJson DefinitionsFetched (list wordDecoder)
        }

-- WordDecoder
wordDecoder : Decoder Word
wordDecoder =
    Json.Decode.map2 Word
        (field "word" string)
        (field "meanings" (list meaningDecoder))

-- MeaningDecoder
meaningDecoder : Decoder Meaning
meaningDecoder =    
    Json.Decode.map2 Meaning
        (field "partOfSpeech" string)
        (field "definitions" (list definitionDecoder))

-- DefinitionDecoder
definitionDecoder : Decoder Definition
definitionDecoder =
    Json.Decode.map Definition
        (field "definition" string)

-- TextToList
textToList : String -> List String
textToList allword =
    String.words allword



