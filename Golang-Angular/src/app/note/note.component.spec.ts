import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Note } from './todo.component';

describe('Note', () => {
  let component: Note;
  let fixture: ComponentFixture<Note>;

  beforeEach(async(() => {
   TestBed.configureTestingModule({
      declarations: [
         Note
      ],
   }).compileComponents();
}));

  it('component should be created', async(() => {
    const fixture = TestBed.createComponent(Note);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });
}));

  it(`component should have title 'app-note'`, async(() => {
     const fixture = TestBed.createComponent(Note);
     const app = fixture.debugElement.componentInstance;
     expect(app.title).toEqual('app-note');
}));

