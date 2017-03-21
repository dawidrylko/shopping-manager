import { Directive, OnInit, Input, ViewContainerRef, TemplateRef } from '@angular/core';

@Directive({
  selector: '[ifRole]'
})
export class IfRoleDirective implements OnInit {

  @Input('ifRole')
  public roles: string[];

  private hasView: boolean = false;

  constructor(
    private viewContainerRef: ViewContainerRef,
    private templateRef: TemplateRef<any>
  ) { }

  ngOnInit() {
    this.updateViewVisibility();
  }

  private updateViewVisibility(): void {
    let isRole = !!this.roles.length;

    if (isRole) {
      this.viewContainerRef.createEmbeddedView(this.templateRef);
    } else {
      this.viewContainerRef.clear();
    }

    this.hasView = isRole;
  }

}
