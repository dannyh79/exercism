module Leap
  ( isLeapYear
  ) where

import Prelude

isLeapYear :: Int -> Boolean
-- isLeapYear year = (year `mod` 4 == 0) && (year `mod` 100 /= 0 || year `mod` 400 == 0)
isLeapYear year
  | year `mod` 400 == 0 = true
  | year `mod` 100 == 0 = false
  | year `mod` 4 == 0 = true
  | otherwise = false
