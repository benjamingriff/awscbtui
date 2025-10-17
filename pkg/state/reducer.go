package state
//
//
//
// func Reduce(s *AppState, a any) {
//     switch act := a.(type) {
//     case FocusNext:
//         s.UI.FocusedView = nextPanel(s.UI.FocusedView)
//
//     case FocusPrev:
//         s.UI.FocusedView = prevPanel(s.UI.FocusedView)
//
//     case FocusView:
//         s.UI.FocusedView = act.View
//
//     case MoveSelection:
//         switch act.View {
//         case ViewProjects:
//             n := len(s.Data.Projects)
//             if n == 0 { break }
//             s.UI.SelectedProject = clampIndex(s.UI.SelectedProject+act.Delta, n)
//         case ViewBuilds:
//             proj := currentProjectName(s)
//             builds := s.Data.Builds[proj]
//             if len(builds) == 0 { break }
//             s.UI.SelectedBuild = clampIndex(s.UI.SelectedBuild+act.Delta, len(builds))
//         }
//     }
// }
//
//
// func Reduce(s *AppState, a any) {
//     switch a.(type) {
//     case FocusNext:
//         s.UI.FocusedView = nextPanel(s.UI.FocusedView)
//     case FocusPrev:
//         s.UI.FocusedView = prevPanel(s.UI.FocusedView)
// 	}
// }
//
// func nextPanel(p View) View {
// 	switch p {
// 	case ViewProjects:
// 		return ViewBuilds
// 	case ViewBuilds:
// 		return ViewLogs
// 	case ViewLogs:
// 		return ViewStatus
// 	case ViewStatus:
// 		return ViewProjects
// 	default:
// 		return ViewProjects
// 	}
// }
//
// func prevPanel(p View) View {
//     switch p {
//     case ViewProjects:
//         return ViewStatus
//     case ViewBuilds:
//         return ViewProjects
//     case ViewLogs:
//         return ViewBuilds
//     case ViewStatus:
//         return ViewLogs
//     default:
//         return ViewProjects
//     }
// }
//
// func currentProjectName(s *AppState) string {
//     if s.UI.SelectedProject >= 0 && s.UI.SelectedProject < len(s.Data.Projects) {
//         return s.Data.Projects[s.UI.SelectedProject].Name
//     }
//     return ""
// }
