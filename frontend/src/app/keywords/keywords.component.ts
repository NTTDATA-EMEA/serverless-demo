import {HttpClient} from '@angular/common/http';
import {Component, OnInit} from '@angular/core';
import {environment} from "../../environments/environment";


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

  public isLoading: boolean = true;

  constructor(private readonly http: HttpClient) {
  }

  onLoad(): void {
    this.isLoading = false;
  }

  ngOnInit(): void {
    this.http.get(this.keywordsUrl).subscribe({
      next: k => {
        this.keywords = Object.keys(k);
      },
      complete: () => {
        this.isLoading = false;
      }
    });
  }

  showBuzzwords(keyword: string) {
    this.isLoading = true;
    this.selectedKeyword = keyword;
    this.buzzwords = [];
    keyword = keyword.startsWith('#') ? keyword.substring(1) : keyword;
    this.http.get(this.buzzwordsUrl + '/' + keyword)
      .subscribe({
        next: (k: any) => {
          this.buzzwords = Object.keys(k.buzzwords);
        },
        error: () => {
          this.buzzwords = [];
          this.isLoading = false;
        },
        complete: () => {
          this.isLoading = false;
        }
      });
  }
}
