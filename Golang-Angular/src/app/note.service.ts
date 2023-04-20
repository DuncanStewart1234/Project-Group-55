import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class NoteService {
  constructor(private httpClient: HttpClient) {}

  getNoteList() {
    return this.httpClient.get(environment.gateway + '/notes');
  }

  addNote(note: Note) {
    return this.httpClient.post(environment.gateway + '/notes', note);
  }

  deleteNote(note: Note) {
    return this.httpClient.delete(environment.gateway + '/notes/' + note.id);
  }
}

export class Note {
  id: string;
  title: string;
  category: string;
  message: string;
}