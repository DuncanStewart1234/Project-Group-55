import { ComponentFixture, TestBed } from '@angular/core/testing';
import { TodoComponent } from './todo.component';

describe('TodoComponent', () => {
  let component: TodoComponent;
  let fixture: ComponentFixture<TodoComponent>;

  beforeEach(async(() => {
   TestBed.configureTestingModule({
      declarations: [
         TodoComponent
      ],
   }).compileComponents();
}));

  it('component should be created', async(() => {
    const fixture = TestBed.createComponent(TodoComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });
}));

  it(`component should have title 'app-todo'`, async(() => {
     const fixture = TestBed.createComponent(TodoComponent);
     const app = fixture.debugElement.componentInstance;
     expect(app.title).toEqual('app-todo');
}));
