import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {

  constructor(private http: HttpClient) { }

  isAuthenticated(): boolean {
    // Check if the user is authenticated (e.g., check a token in local storage)
    return !!localStorage.getItem('userToken');
  }
}
