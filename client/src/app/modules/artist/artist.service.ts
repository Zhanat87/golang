import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';

import {CommonListService} from "../../common/services/list.service";
import {Artist} from "./artist";
import {Observable} from "rxjs";
import { LocalStorageService } from 'angular-2-local-storage';

@Injectable()
export class ArtistService extends CommonListService {

    public url = 'v1/artists';

    constructor (public http: AuthHttp, protected localStorageService: LocalStorageService) {
        super();
    }

    public map(data): Observable<Artist> {
        return data;
    }

    public mapAll(data): Observable<Artist[]> {
        return data;
    }

}