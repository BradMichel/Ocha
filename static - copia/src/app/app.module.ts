import { NgModule }      from '@angular/core';
import { MaterialModule } from '@angular/material';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent }  from './app.component';
import { Toolbar} from './ts/toolbar.component'





export class PizzaPartyAppModule { }
@NgModule({
  imports:      [ BrowserModule,MaterialModule.forRoot() ],
  declarations: [ AppComponent , Toolbar],
  bootstrap:    [ AppComponent , Toolbar]
})
export class AppModule { }
