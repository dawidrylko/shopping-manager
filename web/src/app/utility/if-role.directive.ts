import { Directive, OnInit, Input } from '@angular/core';

@Directive({
  selector: '[ifRole]'
})
export class IfRoleDirective implements OnInit {

  @Input('ifRole')
  public roles: string[];

  constructor() { }

  ngOnInit() {
  }

}
