import { Component } from '@angular/core';
import { DatetimeComponent } from '../datetime/datetime.component';
import { TodoComponent } from '../todo/todo.component';
import { NoteComponent } from '../note/note.component';
import { EventComponent } from '../event/event.component';
import { MapComponent } from '../map/map.component';
import { LoginService } from '../login.service';
import { WeatherComponent } from '../weather/weather.component';
import { Router } from '@angular/router';


@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.css']
})
export class HomepageComponent {
  constructor(private router: Router, private loginService: LoginService){ }

  logOut() {
    this.loginService.logoutUser().subscribe(() => {
        this.router.navigate(["/login"])
    })
}
}
