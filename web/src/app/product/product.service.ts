import { Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';

import { Product } from './product.model';

@Injectable()
export class ProductService {

  private createUrl: string = 'http://localhost:8001/product';
  private listUrl: string = 'http://localhost:8001/product';

  constructor(
    private http: Http
  ) { }

  public create(product: Product){
    let headers = new Headers(
      { 
        'Content-Type': 'application/json'
      }
    );
    let options = new RequestOptions({ headers: headers });
    console.log('options', options)

    return this.http.post(this.createUrl, product, options)
      .map((response: Response) => {
        console.log(response);
        let body = response.json();
    return body.data || { };
      })
      .catch(this.handleError);
  }
  
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
