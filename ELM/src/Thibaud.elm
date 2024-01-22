module Thibaud exposing (..)

import Browser
import Html exposing (..)
import Html.Attributes exposing (style)
import Html.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, map,map2, field, string,list)

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
  | Success (List Word)

type alias Word = 
    { word : String
    , meanings : List Meaning
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
  (Loading, getWord)
  
type Msg
  = GotWord (Result Http.Error (List Word)) 


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotWord result ->
      case result of
        Ok word ->
          (Success word, Cmd.none)

        Err _ ->
          (Failure, Cmd.none)

subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



-- VIEW


view : Model -> Html Msg
view model =
  case model of
    Failure ->
      text "I was unable to load your book."

    Loading ->
      text "Loading..."

    Success fullWords ->
      div []
      [div [] (List.map (\fullWord -> text fullWord.word) fullWords)
      , div [] (List.map viewMeaning fullWords)
      ]

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

getWord : Cmd Msg
getWord = 
  Http.get
      { url = "https://api.dictionaryapi.dev/api/v2/entries/en/think"
      , expect = Http.expectJson GotWord (list wordDecoder)
      }


wordDecoder : Decoder Word
wordDecoder =
    Json.Decode.map2 Word
        (field "word" string)
        (field "meanings" (list meaningDecoder))
        
meaningDecoder : Decoder Meaning
meaningDecoder =
    Json.Decode.map2 Meaning
        (field "partOfSpeech" string)
        (field "definitions" (list definitionDecoder))
        
definitionDecoder : Decoder Definition
definitionDecoder =
    Json.Decode.map Definition
        (field "definition" string)