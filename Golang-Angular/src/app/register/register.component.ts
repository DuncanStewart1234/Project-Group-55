import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})

export class RegisterComponent {
    firstName : string ="";
    lastName : string ="";
    email : string ="";
    username : string ="";
    password : string ="";
    show: boolean= false;

    // constructor(private loginService: LoginService) { }
    constructor(private router: Router){ }
    submit(){
        console.log("user name is " + this.username)
        // this.loginService.loginUser(this.username, this.password);
        this.clear();
    }

    clear(){
    this.firstName ="";
    this.lastName ="";
    this.email ="";
    this.username ="";
    this.password = "";
    this.show = true;
    }

    returnHome(){
        this.router.navigate(['']);
    }
}
