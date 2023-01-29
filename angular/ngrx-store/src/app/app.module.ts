import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { StoreModule } from '@ngrx/store';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MyCounterComponent } from './my-counter/my-counter.component';
import { LoginPageComponent } from './login-page/login-page.component';

import { counterReducer } from '../reducer/counter.reducer';
import { scoreboardReducer } from '../reducer/scoreboard.reducer';

@NgModule({
  declarations: [AppComponent, MyCounterComponent, LoginPageComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    StoreModule.forRoot({ count: counterReducer, game: scoreboardReducer }),
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
