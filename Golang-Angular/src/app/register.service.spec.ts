import { TestBed } from '@angular/core/testing';

import { RegisterService } from './register.service';
import {HttpClientModule} from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

describe('RegisterService', () => {
  let service: RegisterService;

  beforeEach(() => {
    TestBed.configureTestingModule({imports: [HttpClientModule, FormsModule, BrowserModule],
      providers: [RegisterService]});
    service = TestBed.inject(RegisterService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
