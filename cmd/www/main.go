package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, htmlBody)
	})
	http.ListenAndServe(":5000", nil)
}

const htmlBody = `
<!DOCTYPE html>
<html>
  <head>
    <style>
      html, body {
        height: 100%;
        margin: 0;
        padding: 0;
      }
      #map {
        height: 100%;
      }
    </style>
  </head>
  <body>
    <div id="map"></div>
    <script>
      var map;
      
      function initMap() {
        var infowindow = new google.maps.InfoWindow();

        map = new google.maps.Map(document.getElementById('map'), {
          zoom: 13,
          center: new google.maps.LatLng(37.7879, -122.4075)
        });

        google.maps.event.addListener(map, 'click', function() {
          infowindow.close();
        });

        map.data.addListener('click', function(event) {
          infowindow.setContent(event.feature.getProperty('name')+"<br>"+event.feature.getProperty('phone_number'));
          infowindow.setPosition(event.latLng);
          infowindow.setOptions({pixelOffset: new google.maps.Size(0,-34)});
          infowindow.open(map);
        });

        map.data.loadGeoJson('http://35.194.90.189/api?inDate=2015-04-09&outDate=2015-04-10');
      }
    </script>
    <script type="text/javascript" src="http://maps.google.com/maps/api/js?key=AIzaSyDCoNkb_Fea5Wyur5T_1Pya-SZLyOhInlg&sensor=false&callback=initMap" async defer></script>
  </body>
</html>
`
