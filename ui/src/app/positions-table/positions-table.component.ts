import { Component, Output, EventEmitter, Input} from '@angular/core';
import { HttpClient } from '@angular/common/http';

export interface Position {
  id: number;
  title: string;
  salary: number;
  description: string;
}

@Component({
  selector: 'app-positions-table',
  templateUrl: './positions-table.component.html',
  styleUrls: ['./positions-table.component.css']
})
export class PositionsTableComponent {
  positions: Position[] = [];
  @Output() positions_change: EventEmitter<any> = new EventEmitter<any>();

  @Input() worker_position_ids: number[] = [];

  link = 'http://localhost:8080/positions';

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.readData();
  }

  readData() {
    this.http.get<Position[]>(this.link).subscribe(data => {
      this.positions = data;
      const position_ids: number[] = [];
      for (let i = 0; i < this.positions.length; i++) {
        position_ids.push(this.positions[i].id);
      }
      this.positions_change.emit(position_ids);
      this.positions_change.emit(this.positions);
      //sort by id
      this.positions.sort((a, b) => (a.id > b.id) ? 1 : -1);
    });
  }

  editData(position_old: Position, new_data: any) {
    this.http.put<Position>(this.link + '/' + position_old.id, new_data).subscribe(data => {
      this.readData();
    }, error => {
      alert("Щось пішло не так при редагуванні даних");
    });
  }

  deleteData(position: Position) {
    this.http.delete<Position>(this.link + '/' + position.id).subscribe(data => {
      this.readData();
    }, error => {
      alert("Щось пішло не так при видаленні");
    });
  }

  addData(data: any) {
    this.http.post<Position>(this.link, data).subscribe(data => {
      this.readData();
    }, error => {
        alert("Щось пішло не так при додаванні");
      });
  }

  addRowPlaceholder() {
    const element_placeholder = document.getElementById("addRowTable3");
    const element_button = document.getElementById("addRowButton3");

    if (element_placeholder && element_button) {
      element_placeholder.style.display = "";
      element_button.style.display = "none";
    }
  }

  addRow() {
    const element_title = document.getElementById("addRowTitle3") as HTMLInputElement;
    const element_salary = document.getElementById("addRowSalary3") as HTMLInputElement;
    const element_description = document.getElementById("addRowDescription3") as HTMLInputElement;

    const new_title = element_title.value;
    const new_salary = parseInt(element_salary.value);
    const new_description = element_description.value;

    //title (6-100), salary (>0), description (6-1000)
    if (new_title.length < 6 || new_title.length > 100) {
      alert("Назва повинна містити від 6 до 100 символів");
      return;
    }

    if (new_salary <= 0) {
      alert("Зарплата повинна бути більше 0");
      return;
    }

    if (new_description.length < 6 || new_description.length > 1000) {
      alert("Опис повинен містити від 6 до 1000 символів");
      return;
    }

    const new_position = {
      "title": new_title,
      "salary": new_salary,
      "description": new_description
    };

    this.addData(new_position);

    element_title.value = "";
    element_salary.value = "";
    element_description.value = "";
  }

  closeAddRowPlaceholder() {
    const element_placeholder = document.getElementById("addRowTable3");
    const element_button = document.getElementById("addRowButton3");

    if (element_placeholder && element_button) {
      element_placeholder.style.display = "none";
      element_button.style.display = "";
    }

    const element_title = document.getElementById("addRowTitle3") as HTMLInputElement;
    const element_salary = document.getElementById("addRowSalary3") as HTMLInputElement;
    const element_description = document.getElementById("addRowDescription3") as HTMLInputElement;

    element_title.value = "";
    element_salary.value = "";
    element_description.value = "";
  }

  addEditRow(position: Position) {
    const element_editrow = document.getElementById("edit3"+position.id);
    const element_editbuttons = document.getElementById("edit-buttons3"+position.id);

    if (element_editrow && element_editbuttons) {
      element_editrow.style.display = "";
      element_editbuttons.style.display = "none";
    }

    const element_id = document.getElementById("id3"+position.id) as HTMLInputElement;
    const element_title = document.getElementById("title3"+position.id) as HTMLInputElement;
    const element_salary = document.getElementById("salary3"+position.id) as HTMLInputElement;
    const element_description = document.getElementById("description3"+position.id) as HTMLInputElement;

    element_id.value = position.id.toString();
    element_title.value = position.title;
    element_salary.value = position.salary.toString();
    element_description.value = position.description;

    for (let i = 0; i < this.positions.length; i++) {
      if (this.positions[i].id != position.id) {
        const element_editbuttons = document.getElementById("edit-buttons3"+this.positions[i].id);
        if (element_editbuttons){
          element_editbuttons.style.display = "none";
        }
      }
    }
  }

  deleteRow(position: Position) {
    if (this.worker_position_ids.includes(position.id)) {
      alert("Ця посада використовується");
      return;
    }
    this.deleteData(position);
  }

  approveEditRow(position: Position) {
    const element_id = document.getElementById("id3"+position.id) as HTMLInputElement;
    const element_title = document.getElementById("title3"+position.id) as HTMLInputElement;
    const element_salary = document.getElementById("salary3"+position.id) as HTMLInputElement;
    const element_description = document.getElementById("description3"+position.id) as HTMLInputElement;

    const new_id = parseInt(element_id.value);
    const new_title = element_title.value;
    const new_salary = parseInt(element_salary.value);
    const new_description = element_description.value;

    // title (6-100), salary (>0), description (6-1000)
    if (new_title.length < 6 || new_title.length > 100) {
      alert("Назва повинна містити від 6 до 100 символів");
      return;
    }

    if (new_salary <= 0) {
      alert("Зарплата повинна бути більше 0");
      return;
    }
    
    if (new_description.length < 6 || new_description.length > 1000) {
      alert("Опис повинен містити від 6 до 1000 символів");
      return;
    }

    if (new_title==position.title && new_salary==position.salary && new_description==position.description) {
      alert("Дані не змінилися");
      return;
    }

    const new_data = {
      "title": new_title,
      "salary": new_salary,
      "description": new_description
    };

    this.editData(position, new_data);
    this.closeEditRow(position);
  }

  closeEditRow(position: Position) {
    const element_editrow = document.getElementById("edit3"+position.id);
    const element_editbuttons = document.getElementById("edit-buttons3"+position.id);

    if (element_editrow && element_editbuttons) {
      element_editrow.style.display = "none";
      element_editbuttons.style.display = "";
    }

    const element_id = document.getElementById("id3"+position.id) as HTMLInputElement;
    const element_title = document.getElementById("title3"+position.id) as HTMLInputElement;
    const element_salary = document.getElementById("salary3"+position.id) as HTMLInputElement;
    const element_description = document.getElementById("description3"+position.id) as HTMLInputElement;

    element_id.value = "";
    element_title.value = "";
    element_salary.value = "";
    element_description.value = "";

    for (let i = 0; i < this.positions.length; i++) {
      if (this.positions[i].id != position.id) {
        const element_editbuttons = document.getElementById("edit-buttons3"+this.positions[i].id);
        if (element_editbuttons){
          element_editbuttons.style.display = "";
        }
      }
    }
  }
}
