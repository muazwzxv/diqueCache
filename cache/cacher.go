package cache

import "time"

/**
  Might do 
  - make the transactions atomic
  - implement some consensus protocol
*/

type ICacher interface {
  Set([]byte, []byte, time.Duration) error 
  Has([]byte) bool
  Get([]byte) ([]byte, error)
  Delete([]byte) error
}
