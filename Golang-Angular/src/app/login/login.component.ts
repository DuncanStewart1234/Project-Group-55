import { Component } from '@angular/core';
import { LoginService, UserLogin } from '../login.service';
import { Router } from '@angular/router';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.css']
  })

export class LoginComponent {
    username : string ="";
    password : string ="";
    show: boolean= false;

    // constructor(private loginService: LoginService) { }
    constructor(private router: Router, private loginService: LoginService){ }

    submit(){
        console.log("user name is " + this.username)
        this.loginService.loginUser(this.username, this.password).subscribe(() => {
            this.clear();
        });
    }

    clear(){
    this.username ="";
    this.password = "";
    this.show = true;
    }

    returnHome(){
        this.router.navigate(['']);
    }
}
