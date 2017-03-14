import { ShoppingManagerPage } from './app.po';

describe('shopping-manager App', () => {
  let page: ShoppingManagerPage;

  beforeEach(() => {
    page = new ShoppingManagerPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
