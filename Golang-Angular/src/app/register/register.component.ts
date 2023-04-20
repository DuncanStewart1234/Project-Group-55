import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { RegisterService, UserRegister } from '../register.service';
import { LoginService } from '../login.service';


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
    constructor(private router: Router, private registerService: RegisterService, private loginService: LoginService){ }
    // constructor(private router: Router){ }

    submit(){
        console.log("user name is " + this.username)
        var newUser : UserRegister = {
            first: this.firstName,
            last: this.lastName,
            email: this.email,
            uname: this.username,
            password: this.password
        }
        // this.registerService.registerUser(newUser)
        this.registerService.registerUser(newUser).subscribe(() => {
            this.loginService.loginUser(newUser.uname, newUser.password)
        });
        // this.loginService.loginUser(this.username, this.password);
        this.clear();
    }

    clear() {
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
