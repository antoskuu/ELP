module Thibaud exposing (..)

import Browser
import Html exposing (..)
import Html.Attributes exposing (style)
import Html.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, field, string, list, map2, map)
import Browser.Navigation exposing (load)
import Random exposing (step, Generator)


main =
    Browser.element
        { init = init
        , update = update
        , subscriptions = subscriptions
        , view = view
        }

type Model
    = Failure
    | Loading
    | Success SelectedWord (List Meaning)

type alias SelectedWord =
    { word : String
    , meanings : List Meaning
    , wordSelected : Bool
    }

type alias Meaning =
    { partOfSpeech : String
    , definitions : List Definition
    }

type alias Definition =
    { definition : String
    }

init : () -> (Model, Cmd Msg)
init _ =
    (Loading, Http.get { url = "/thousand_words_things_explainer.txt", expect = Http.expectString WordFetched })

type Msg
    = FetchWord
    | WordFetched (Result Http.Error String)
    | FetchDefinitions String
    | DefinitionsFetched (Result Http.Error (List Definition))

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        FetchWord ->
            (Loading, Http.get { url = "/thousand_words_things_explainer.txt", expect = Http.expectString WordFetched })

        WordFetched (Ok wordText) ->
            let
                wordList =
                    textToList wordText
                randomWord =
                    getRandomWord wordList
            in
            (Success { word = randomWord, meanings = [], wordSelected = False }, getDefinitionsCmd randomWord)

        WordFetched (Err _) ->
            (Failure, Cmd.none)

        FetchDefinitions selectedWord ->
            (Loading, getDefinitionsCmd selectedWord)

        DefinitionsFetched (Ok definitions) ->
            case model of
                Success selectedWord _ ->
                    (Success { selectedWord | meanings = definitions, wordSelected = True }, Cmd.none)

                _ ->
                    (Failure, Cmd.none)

        DefinitionsFetched (Err _) ->
            (Failure, Cmd.none)

subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none

view : Model -> Html Msg
view model =
    case model of
        Failure ->
            text "I was unable to load your book."

        Loading ->
            text "Loading..."

        Success fullWords ->
            div []
                [ div [] [ text "Word: " ]
                , div [] [ text fullWords.word ]
                , div [] (List.map viewMeaning fullWords.meanings)
                ]

viewMeaning : Meaning -> Html Msg
viewMeaning meaning =
    div []
        [ div [] [ text ("Part of Speech: " ++ meaning.partOfSpeech) ]
        , div [] (List.map viewDefinition meaning.definitions)
        ]

viewDefinition : Definition -> Html Msg
viewDefinition definition =
    div []
        [ div [] [ text ("Definition: " ++ definition.definition) ]
        ]

getDefinitionsCmd : String -> Cmd Msg
getDefinitionsCmd word =
    Http.get
        { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ word
        , expect = Http.expectJson DefinitionsFetched (list definitionDecoder)
        }

wordDecoder : Decoder Meaning
wordDecoder =
    Json.Decode.map2 Meaning
        (field "partOfSpeech" string)
        (field "definitions" (list definitionDecoder))

meaningDecoder : Decoder Meaning
meaningDecoder =
    Json.Decode.map2 Meaning
        (field "partOfSpeech" string)
        (field "definitions" (list definitionDecoder))

definitionDecoder : Decoder Definition
definitionDecoder =
    Json.Decode.map Definition
        (field "definition" string)


textToList : String -> List String
textToList allword =
    String.words allword



