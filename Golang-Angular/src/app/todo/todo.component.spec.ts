import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Todo } from './todo.component';

describe('TodoComponent', () => {
  let component: Todo;
  let fixture: ComponentFixture<Todo>;

  beforeEach(async(() => {
   TestBed.configureTestingModule({
      declarations: [
         Todo
      ],
   }).compileComponents();
}));

  it('component should be created', async(() => {
    const fixture = TestBed.createComponent(Todo);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });
});
