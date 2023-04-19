import { Component } from '@angular/core';
import { LeafletDirective } from '@asymmetrik/ngx-leaflet';
import * as Leaflet from 'leaflet'; 
import { GeoSearchControl, OpenStreetMapProvider } from 'leaflet-geosearch';
// import {} from 'googlemaps';

Leaflet.Icon.Default.imagePath = 'assets/';

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css']
})
export class MapComponent {
  map!: Leaflet.Map;
  markers: Leaflet.Marker[] = [];
  lat: number;
  lng: number;
  index: number;
  options = {
	attributionControl: false,
    layers: [
      Leaflet.tileLayer('https://{s}.basemaps.cartocdn.com/dark_all/{z}/{x}/{y}{r}.png'),
      // Leaflet.tileLayer('http://{s}.tile.osm.org/{z}/{x}/{y}.png')
    ],
    zoom: 16,
    center: { lat: 29.643946, lng: -82.355659 }
  }


  initMarkers() {
    const initialMarkers = [
      {
        position: { lat: 29.6499, lng: -82.3487 },
        draggable: true,
        title: "test",
      },
      {
        position: { lat: 29.6481, lng: -82.3437 },
        draggable: false
      },
      {
        position: { lat: 29.643946, lng: -82.355659 },
        draggable: true
      }
    ];
    for (this.index = 0; this.index < initialMarkers.length; this.index++) {
      const data = initialMarkers[this.index];
      const marker = this.generateMarker(data, this.index);
      // marker.addTo(this.map).bindPopup(`<b>${data.position.lat},  ${data.position.lng}</b>`);
      marker.addTo(this.map).bindPopup(`<b>${data.position.lat}</b>`);
      // marker.bindTooltip("Marston Science Library", 
      // {
      //     permanent: true,
      //     direction: 'top'
      // }
      // );
      this.map.panTo(data.position);
      this.markers.push(marker)
    }

  }

  createMarker(name: string, lat: number, lng: number) {
    const data = {
      position: { lat: lat, lng: lng },
      draggable: true
    };

    const marker = this.generateMarker(data, this.index++);

    marker.addTo(this.map).bindPopup(`<b>${name}</b>`);
    this.map.panTo(data.position);
    this.markers.push(marker);

  }

  generateMarker(data: any, index: number) {
    return Leaflet.marker(data.position, { draggable: data.draggable })
      .on('click', (event) => this.markerClicked(event, index))
      .on('dragend', (event) => this.markerDragEnd(event, index));
  }

  onMapReady($event: Leaflet.Map) {
    this.map = $event;
    // Leaflet.Control;
    const provider = new OpenStreetMapProvider();

    const searchControl =  GeoSearchControl({
      provider: provider,
      position: 'topright',
      style: 'button',
      autoCompleteDelay: 100,

    });

    this.map.addControl(searchControl);

    // this.initMarkers();                                                       // INIT MARKERS
    // Leaflet.tileLayer('https://{s}.tile.osm.org/{z}/{x}/{y}.png', {  attribution: '&copy; <a href="https://osm.org/copyright">OpenStreetMap</a> contributors'}).addTo(this.map);
    // Leaflet.Control.geocoder().addTo(this.map);
    
  }

  mapClicked($event: any) {
    console.log($event.latlng.lat, $event.latlng.lng);
    // this.markers = 
    this.lat = $event.latlng.lat;
    this.lng = $event.latlng.lng;
    this.createMarker("Marker", $event.latlng.lat, $event.latlng.lng);
  }

  markerClicked($event: any, index: number) {
    console.log($event.latlng.lat, $event.latlng.lng);
    this.lat = $event.latlng.lat;
    this.lng = $event.latlng.lng;
  }

  markerDragEnd($event: any, index: number) {
    console.log($event.target.getLatLng());
  } 

  testFunc(){
    alert('test');
  }
}
