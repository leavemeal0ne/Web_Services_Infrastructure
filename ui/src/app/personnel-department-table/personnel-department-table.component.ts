import { Component, Input } from '@angular/core';
import { Worker } from '../workers-table/workers-table.component';
import { Position } from '../positions-table/positions-table.component';

interface TableRow {
  id_worker: number;
  full_name: string;
  id_position: number;
  title: string;
  salary: number;
  description: string;
}

@Component({
  selector: 'app-personnel-department-table',
  templateUrl: './personnel-department-table.component.html',
  styleUrls: ['./personnel-department-table.component.css']
})
export class PersonnelDepartmentTableComponent {
    @Input() workers: Worker[] = [];
    @Input() positions: Position[] = [];

    table_rows: TableRow[] = [];

    ngOnInit() {
    }

}
