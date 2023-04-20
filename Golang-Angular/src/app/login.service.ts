import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class LoginService {
  constructor(private httpClient: HttpClient) {}

  loginPage() {
    return this.httpClient.get(environment.gateway + '/login')
  }

  loginUser(username: string, password: string) {
    return this.httpClient.post(environment.gateway + '/login', { username, password })
  }

}

export class UserLogin {
  uname: string;
  passwd: string;
}