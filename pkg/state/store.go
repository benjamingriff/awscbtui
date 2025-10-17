package state
//
// import "sync"
//
// type Store struct {
//     mu sync.Mutex
//     s  *AppState
// }
//
// func NewStore(s *AppState) *Store {
// 	return &Store{s: s} 
// }
//
// func (st *Store) Dispatch(a any) {
//     st.mu.Lock()
//     Reduce(st.s, a)
//     st.mu.Unlock()
// }
//
// func (st *Store) Snapshot() AppState {
//     st.mu.Lock()
//     defer st.mu.Unlock()
//     // Return a copy for safe reads. Implement a shallow copy method if needed.
//     cp := *st.s
//     return cp
// }
//
// func (st *Store) StateUnsafe() *AppState {
//     // Only use on the UI thread if you guarantee no concurrent writes.
//     return st.s
// }
