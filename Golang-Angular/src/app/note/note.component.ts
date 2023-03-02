import { Component, OnInit } from '@angular/core';
import { NoteService, Note } from '../note.service';
@Component({
  selector: 'app-note',
  templateUrl: './note.component.html',
  styleUrls: ['./note.component.css']
})

export class NoteComponent {

  activeNotes: Note[];
  noteMessage: string;

  constructor(private noteService: NoteService) { }

  ngOnInit() {
    this.getAll();
  }

  getAll() {
    this.noteService.getNoteList().subscribe((data: any) => 
    {
      this.activeNotes = data;
    });
  }

  addNote() {
    // console.log(this.noteMessage);
    var newNote : Note = {
      message: this.noteMessage,
      id: '',
    };

    this.noteService.addNote(newNote).subscribe(() => {
      this.getAll();
      this.noteMessage = '';
    });
  }

  deleteNote(note: Note) {
    this.noteService.deleteNote(note).subscribe(() => {
      this.getAll();
    })
  }
}
