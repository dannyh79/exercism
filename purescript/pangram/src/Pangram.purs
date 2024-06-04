module Pangram
  ( isPangram
  ) where

import Prelude

import Data.Set (fromFoldable, subset)
import Data.String (toLower)
import Data.String.CodeUnits (toCharArray)

isPangram :: String -> Boolean
isPangram str = subset alphabetSet strSet
  where
  alphabetSet = fromFoldable $ toCharArray "abcdefghijklmnopqrstuvwxyz"
  strSet = fromFoldable $ toCharArray $ toLower str
