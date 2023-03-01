import { ComponentFixture, TestBed } from '@angular/core/testing';
import { NoteComponent } from './note.component';

describe('NoteComponent', () => {
  let component: NoteComponent;
  let fixture: ComponentFixture<NoteComponent>;

  beforeEach(async(() => {
   TestBed.configureTestingModule({
      declarations: [
         NoteComponent
      ],
   }).compileComponents();
}));

  it('component should be created', async(() => {
    const fixture = TestBed.createComponent(NoteComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });
}));

  it(`component should have title 'app-note'`, async(() => {
     const fixture = TestBed.createComponent(NoteComponent);
     const app = fixture.debugElement.componentInstance;
     expect(app.title).toEqual('app-note');
}));

