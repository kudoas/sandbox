import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Observable, lastValueFrom } from 'rxjs';
import { HousingLocation } from '../housinglocation';

@Component({
  selector: 'app-sample-search',
  standalone: true,
  imports: [CommonModule, HttpClientModule],
  template: `
    <form action="">
      <input type="text" placeholder="Search" #search />
      <button type="button" (click)="onSearch(search.value)">Search</button>
    </form>

    <ul *ngFor="let housingLocation of HousingLocations">
      <li>{{ housingLocation.name }}</li>
    </ul>
  `,
  styleUrls: ['./sample-search.component.scss'],
})
export class SampleSearchComponent {
  url = 'http://localhost:3000/locations';
  http = inject(HttpClient);
  HousingLocations: HousingLocation[] = [];

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

  getAllHousingLocations(keyword: string): Observable<HousingLocation[]> {
    return this.http
      .get(`${this.url}?keyword=${keyword}`)
      .pipe((data) => data as Observable<HousingLocation[]>);
  }
}
