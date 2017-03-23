import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';

import { Product } from './product.model';

@Injectable()
export class ProductService {

  private listUrl: string = 'http://localhost:8001/product';

  constructor(
    private http: Http
  ) { }
  
  public list(): Observable<Product[]> {
    return this.http.get(this.listUrl)
      .map((response: Response) => {
        return response.json();
      })
      .catch(this.handleError);
  }

  private handleError (error: Response | any) {
    let errMsg: string;

    if (error instanceof Response) {
      const body: any = error.json() || '';
      const err = body.error || JSON.stringify(body);
      errMsg = `${error.status} - ${error.statusText || ''} ${err}`;
    } else {
      errMsg = error.message ? error.message : error.toString();
    }

    console.error(errMsg);
    return Observable.throw(errMsg);
  }

}
