import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Route } from '@angular/router';
import { DashboardComponent } from './dashboard/dashboard.component';
import { SidebarComponent } from './components/sidebar/sidebar.component';
import { MainComponent } from './components/main/main.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { BreadcrumbComponent } from './components/breadcrumb/breadcrumb.component';

export const dashboardRoutes: Route[] = [
  {
    path: '',
    component: DashboardComponent,
    children: [
      {
        path: '',
        loadChildren: () =>
          import('@looped-cms/pages/home').then((m) => m.PagesHomeModule),
      },
      {
        path: 'settings',
        loadChildren: () =>
          import('@looped-cms/pages/settings').then(
            (m) => m.PagesSettingsModule
          ),
      },
    ],
  },
];

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forChild(dashboardRoutes),
    FontAwesomeModule,
  ],
  declarations: [
    DashboardComponent,
    SidebarComponent,
    MainComponent,
    BreadcrumbComponent,
  ],
})
export class DashboardModule {}
