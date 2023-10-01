import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Observable, lastValueFrom } from 'rxjs';
import {
  FormsModule,
  ReactiveFormsModule,
  FormGroup,
  FormControl,
} from '@angular/forms';

import { HousingLocation } from '../housinglocation';

@Component({
  selector: 'app-sample-search',
  standalone: true,
  imports: [CommonModule, HttpClientModule, FormsModule, ReactiveFormsModule],
  template: `
    <form>
      <input
        type="text"
        placeholder="Search"
        [(ngModel)]="keyword"
        [ngModelOptions]="{ standalone: true }"
      />
      <button type="button" (click)="onSearch(keyword)">Search</button>
    </form>
    <form [formGroup]="searchFormGroup">
      <input type="text" placeholder="Search" formControlName="keyword" />
      <button type="button" (click)="onSubmit()">Search</button>
    </form>

    <ul *ngFor="let housingLocation of HousingLocations">
      <li>{{ housingLocation.name }}</li>
    </ul>
  `,
  styleUrls: ['./sample-search.component.scss'],
})
export default class SampleSearchComponent {
  url = 'http://localhost:3000/locations';
  private readonly http = inject(HttpClient);
  HousingLocations: HousingLocation[] = [];
  keyword = '';
  searchFormGroup = new FormGroup({
    keyword: new FormControl(''),
  });

  constructor() {
    lastValueFrom(this.getAllHousingLocations(''))
      .then((data) => {
        this.HousingLocations = data;
      })
      .catch((err) => {
        console.error(err);
      });
  }

  async onSearch(keyword: string) {
    this.HousingLocations = await lastValueFrom(
      this.getAllHousingLocations(keyword)
    );
  }

  async onSubmit() {
    this.HousingLocations = [];
    this.HousingLocations = await lastValueFrom(
      this.getAllHousingLocations(this.searchFormGroup.value.keyword!)
    );
  }

  getAllHousingLocations(keyword: string): Observable<HousingLocation[]> {
    return this.http
      .get(`${this.url}?keyword=${keyword}`)
      .pipe((data) => data as Observable<HousingLocation[]>);
  }
}
