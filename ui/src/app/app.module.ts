import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ClientsTableComponent } from './clients-table/clients-table.component';
import { HttpClientModule } from '@angular/common/http';
import { WorkersTableComponent } from './workers-table/workers-table.component';
import { PositionsTableComponent } from './positions-table/positions-table.component';
import { PersonnelDepartmentTableComponent } from './personnel-department-table/personnel-department-table.component';

@NgModule({
  declarations: [
    AppComponent,
    ClientsTableComponent,
    WorkersTableComponent,
    PositionsTableComponent,
    PersonnelDepartmentTableComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
