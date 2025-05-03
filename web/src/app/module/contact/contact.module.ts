import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';



@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    RouterModule.forChild([
      { path: 'contact' },
      { path: 'contact/create' },
      { path: 'contact/edit/:id' },
      { path: 'company' },
      { path: 'company/create' },
      { path: 'company/edit/:id' },
    ]),
  ]
})
export class ContactModule { }
