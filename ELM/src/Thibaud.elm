module Thibaud exposing (..)

import Array exposing (fromList, get)
import Browser
import Html exposing (Html, div, text, input, button, h1, p, blockquote,ul,li,ol)
import Html.Attributes exposing (placeholder, value, disabled, class, type_, checked)
import Html.Events exposing (onClick, onInput)
import Http exposing (Error)
import Random exposing (int)
import Json.Decode exposing (Decoder, string, list, field, map2, map3)


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
    , randomword : String
    , state : State
    , select : String
    }

type State 
    = Failure
    | Loading
    | Success SelectedWord


type alias SelectedWord =
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
    | FetchDefinitions String
    | DefinitionsFetched (Result Http.Error (List Definition))

-- Init function
init : () -> (Model, Cmd Msg)
init _ =
    ({state = Loading , allWord = []}, Http.get { url = "/thousand_words_things_explainer.txt", expect = Http.expectString WordFetched })

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
                ({ model | allWord = wordList}, RandomNumber index)

        WordFetched (Err _) ->
            (Failure, Cmd.none)

        RandomNumber ind ->
            let 
                selected = get ind (fromList model.allWord)
            in 
            ( {model | randomword = selected, getDefinitionsCmd = getDefinitionsCmd model.randomword })

        DefinitionsFetched (Ok definitions) ->
            ({model | state = Success {word = model.select, meanings = definitions }}, Cmd.none)
                    
        DefinitionsFetched (Err _) ->
            -- En cas d'erreur lors de la récupération des définitions, on met à jour l'état à Failure
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

        Success selectedWord ->
            div []
                [ div [] [ text "Mot: " ]
                , div [] [ text selectedWord.word ]
                , div [] (List.map viewMeaning selectedWord.meanings)
                ]

-- Fonction auxiliaire pour afficher une signification (meaning)
viewMeaning : Meaning -> Html Msg
viewMeaning meaning =
    div []
        [ div [] [ text ("Partie du discours : " ++ meaning.partOfSpeech) ]
        , div [] (List.map viewDefinition meaning.definitions)
        ]

-- Fonction auxiliaire pour afficher une définition (definition)
viewDefinition : Definition -> Html Msg
viewDefinition definition =
    div []
        [ div [] [ text ("Définition : " ++ definition.definition) ]
        ]

-- GetDefinitionsCmd
getDefinitionsCmd : String -> Cmd Msg
getDefinitionsCmd word =
    Http.get
        { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ word
        , expect = Http.expectJson DefinitionsFetched (list definitionDecoder)
        }

-- WordDecoder
wordDecoder : Decoder Meaning
wordDecoder =
    Json.Decode.map2 Meaning
        (field "partOfSpeech" string)
        (field "definitions" (list definitionDecoder))

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

-- GetRandomWord

