import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class RegisterService {
  constructor(private httpClient: HttpClient) {}

  registerPage() {
    return this.httpClient.get(environment.gateway + '/signup');
  }

  registerUser(user: UserRegister) {
    return this.httpClient.post(environment.gateway + '/signup', user);
  }
}

export class UserRegister {
	first: string;
	last: string;
	email: string;
	uname: string;
	password: string;
}