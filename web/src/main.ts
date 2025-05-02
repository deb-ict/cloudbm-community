import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';

import { VERSION } from '@angular/material/core';

console.info(`Angular Material Version: ${VERSION.full}`);

bootstrapApplication(AppComponent, appConfig)
  .catch((err) => console.error(err));
