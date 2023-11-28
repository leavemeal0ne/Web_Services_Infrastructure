import { Component, Output, EventEmitter } from '@angular/core';
import { Worker } from './workers-table/workers-table.component';
import { Position } from './positions-table/positions-table.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Lab4';
  position_ids: number[] = [];
  worker_position_ids: number[] = [];

  positions: Position[] = [];
  workers: Worker[] = [];

  onPositions(data: any) {
    if (typeof data[0] === 'number') {
      this.position_ids = data;
    }
    else {
      this.positions = data;
    }
  }

  onWorkers(data: any) {
    if (typeof data[0] === 'number') {
      this.worker_position_ids = data;
    }
    else {
      this.workers = data;
    }
  }
}
