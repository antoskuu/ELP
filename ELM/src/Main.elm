module Main exposing (..)

import Browser
import Html exposing (Html, text, pre)
import Http
import Array exposing (Array)
import Random exposing (Generator, initialSeed, step, int)
import Maybe exposing (withDefault)


-- MAIN


main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }



-- MODEL


type Model
  = Failure
  | Loading
  | Success (Array String)


init : () -> (Model, Cmd Msg)
init _ =
  (Loading, Http.get { url = "/thousand_words_things_explainer.txt", expect = Http.expectString GotText })



-- UPDATE


type Msg
  = GotText (Result Http.Error String)


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotText result ->
      case result of
        Ok fullText ->
          let
            modifiedText = String.words fullText
            wordArray = Array.fromList modifiedText
          in
          (Success wordArray, Cmd.none)

        Err _ ->
          (Failure, Cmd.none)



-- SUBSCRIPTIONS


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

    Success wordArray ->
      let
        randomWord = getRandomWord wordArray (Random.step (initialSeed 42))
      in
      pre [] [ text randomWord ]



getRandomWord : Array String -> Generator Int -> String
getRandomWord wordArray generator =
  let
    index = withDefault 0 (Random.generate int generator)
  in
  Array.get index wordArray |> withDefault "No word found"