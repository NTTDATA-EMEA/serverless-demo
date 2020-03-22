import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { environment } from "../../environments/environment";


@Component({
  selector: 'app-keywords',
  templateUrl: './keywords.component.html',
  styleUrls: ['./keywords.component.sass']
})
export class KeywordsComponent implements OnInit {

  private keywordsUrl = environment.baseApiUrl + '/state';
  private buzzwordsUrl = environment.baseApiUrl + '/aggregate';

  public keywords: string[] = [];
  public buzzwords: string[] = [];

  public selectedKeyword = '';

  constructor(private readonly http: HttpClient) { }

  ngOnInit(): void {
    this.http.get(this.keywordsUrl).subscribe(k => {
      this.keywords = Object.keys(k);
    });
  }

  showBuzzwords(keyword: string) {
    this.selectedKeyword = keyword;
    this.buzzwords = [];
    keyword = keyword.startsWith('#') ? keyword.substring(1) : keyword;
    this.http.get(this.buzzwordsUrl + '/' + keyword).subscribe((k: any) => {
      this.buzzwords = Object.keys(k.buzzwords);
    });
  }
}
