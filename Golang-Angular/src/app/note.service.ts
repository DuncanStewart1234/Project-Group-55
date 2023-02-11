import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class NoteService {
  constructor(private httpClient: HttpClient) {}

  getNoteList() {
    return this.httpClient.get(environment.gateway + '/note');
  }

  addNote(note: Note) {
    return this.httpClient.post(environment.gateway + '/note', note);
  }


  deleteNote(note: Note) {
    return this.httpClient.delete(environment.gateway + '/note/' + note.id);
  }
}

export class Note {
  id: string;
  message: string;
}