import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: '',
    loadChildren: () => import('./pages/menu/menu.module').then( m => m.MenuPageModule)
  },
  {
    path: 'add-text-event',
    loadChildren: () => import('./dialogs/add-text-event/add-text-event.module').then( m => m.AddTextEventPageModule)
  },
  {
    path: 'update-text-event',
    loadChildren: () => import('./dialogs/update-text-event/update-text-event.module').then( m => m.UpdateTextEventPageModule)
  },
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, { preloadingStrategy: PreloadAllModules })
  ],
  exports: [RouterModule]
})
export class AppRoutingModule { }
