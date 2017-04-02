import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Product } from './../product.model'

@Component({
  selector: 'app-product-form',
  templateUrl: './product-form.component.html',
  styleUrls: ['./product-form.component.scss']
})
export class ProductFormComponent implements OnInit {
  public product: FormGroup;

  constructor(
    private formBuilder: FormBuilder
  ) { }

  ngOnInit() {
    this.product = this.formBuilder.group({
      name: ['Kolanka ozdobne Nr 33 Makaron Lubella'],
      ean: ['5900049003060'],
      manufacturer: ['Lubella'],
      energy: ['351'],
      netWeight: ['500'],
      unitOfMeasure: ['gram'],
      price: [3.59]
    });
  }

  onSubmit({
    value,
    valid
  }) {
    console.log('value', value);
    console.log('valid', valid);
  }

}
