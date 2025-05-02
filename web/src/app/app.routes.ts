import { Routes } from '@angular/router';
import { AuthGuard } from './module/auth/auth.guard';
import { HomeComponent } from './home/home.component';
import { AboutComponent } from './about/about.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';

import { LoginComponent } from './module/auth/page/login/login.component';

export const routes: Routes = [
    { path: 'dashboard', component: HomeComponent, canActivate: [AuthGuard], title: 'Dashboard' },
    { path: 'login', component: LoginComponent },
    //{ path: 'auth', loadChildren: () => import('./module/auth/auth.module').then(m => m.AuthModule) },
    { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
    { path: '**', component: PageNotFoundComponent }
];
