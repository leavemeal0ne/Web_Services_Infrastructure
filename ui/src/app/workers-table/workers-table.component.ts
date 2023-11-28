import { Component, Input, Output, EventEmitter } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export interface Worker {
  id: number;
  full_name: string;
  age: number;
  sex: string;
  position_id: number;
}

@Component({
  selector: 'app-workers-table',
  templateUrl: './workers-table.component.html',
  styleUrls: ['./workers-table.component.css']
})
export class WorkersTableComponent {
  @Input() position_ids: number[] = [];

  @Output() workers_change: EventEmitter<any> = new EventEmitter<any>();
  workers: Worker[] = [];

  link = 'http://localhost:8080/workers';

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.readData();
  }

  readData() {
    this.http.get<Worker[]>(this.link).subscribe(data => {
      this.workers = data;
      const worker_position_ids: number[] = [];
      for (let i = 0; i < this.workers.length; i++) {
        worker_position_ids.push(this.workers[i].position_id);
      }
      this.workers_change.emit(worker_position_ids);
      this.workers_change.emit(this.workers);
      //sort by id
      this.workers.sort((a, b) => (a.id > b.id) ? 1 : -1);
    });
  }

  editData(worker_old: Worker, new_data: any) {
    this.http.put<Worker>(this.link + '/' + worker_old.id, new_data).subscribe(data => {
      this.readData();
    }, error => {
      alert("Щось пішло не так при редагуванні даних");
    });
  }

  deleteData(worker: Worker) {
    this.http.delete<Worker>(this.link + '/' + worker.id).subscribe(data => {
      this.readData();
    }, error => {
      alert("Щось пішло не так при видаленні");
    });
  }

  addData(data: any) {
    this.http.post<Worker>(this.link, data).subscribe(data => {
      this.readData();
    }, error => {
        alert("Щось пішло не так при додаванні");
      });
  }

  addRowPlaceholder() {
    const element_placeholder = document.getElementById("addRowTable2");
    const element_button = document.getElementById("addRowButton2");

    if (element_placeholder && element_button) {
      element_placeholder.style.display = "";
      element_button.style.display = "none";
    }
  }

  addRow() {
    const element_full_name = document.getElementById("addRowFullName2") as HTMLInputElement;
    const element_age = document.getElementById("addRowAge2") as HTMLInputElement;
    const element_sex = document.getElementById("addRowSex2") as HTMLSelectElement;
    const element_position_id = document.getElementById("addRowPositionId2") as HTMLInputElement;

    const new_name = element_full_name.value;
    const new_age = parseInt(element_age.value);
    const new_sex = element_sex.value;
    const new_position_id = parseInt(element_position_id.value);

    //full_name (6-100 chrs), age (18-120), sex ('female','male')
    if (new_name.length < 6 || new_name.length > 100) {
      alert("Ім'я повинно містити від 6 до 100 символів");
      return;
    }

    if (new_age < 18 || new_age > 120) {
      alert("Вік повинен бути від 18 до 120 років");
      return;
    }

    if (this.position_ids.indexOf(new_position_id) == -1) {
      alert("Такої посади не існує");
      return;
    }

    const new_worker = {
      "full_name": new_name, 
      "age": new_age, 
      "sex": new_sex,
      "position_id": new_position_id
    };
    this.addData(new_worker);

    element_full_name.value = "";
    element_age.value = "";
    element_sex.value = "male";
    element_position_id.value = "";
  }

  closeAddRowPlaceholder() {
    const element_placeholder = document.getElementById("addRowTable2");
    const element_button = document.getElementById("addRowButton2");

    if (element_placeholder && element_button) {
      element_placeholder.style.display = "none";
      element_button.style.display = "";
    }

    const element_full_name = document.getElementById("addRowFullName2") as HTMLInputElement;
    const element_age = document.getElementById("addRowAge2") as HTMLInputElement;
    const element_sex = document.getElementById("addRowSex2") as HTMLSelectElement;
    const element_position_id = document.getElementById("addRowPositionId2") as HTMLInputElement;

    element_full_name.value = "";
    element_age.value = "";
    element_sex.value = "male";
    element_position_id.value = "";
  }

  addEditRow(worker: Worker) {
    const element_editrow = document.getElementById("edit2"+worker.id);
    const element_editbuttons = document.getElementById("edit-buttons2"+worker.id);

    if (element_editrow && element_editbuttons) {
      element_editrow.style.display = "";
      element_editbuttons.style.display = "none";
    }

    const element_id = document.getElementById("id2"+worker.id) as HTMLInputElement;
    const element_full_name = document.getElementById("full_name2"+worker.id) as HTMLInputElement;
    const element_age = document.getElementById("age2"+worker.id) as HTMLInputElement;
    const element_sex = document.getElementById("sex2"+worker.id) as HTMLInputElement;
    const element_position_id = document.getElementById("position_id2"+worker.id) as HTMLInputElement;

    element_id.value = worker.id.toString();
    element_full_name.value = worker.full_name;
    element_age.value = worker.age.toString();
    element_sex.value = worker.sex;
    element_position_id.value = worker.position_id.toString();

    for (let i = 0; i < this.workers.length; i++) {
      if (this.workers[i].id != worker.id) {
        const element_editbuttons = document.getElementById("edit-buttons2"+this.workers[i].id);
        if (element_editbuttons){
          element_editbuttons.style.display = "none";
        }
      }
    }
  }

  deleteRow(worker: Worker) {
    this.deleteData(worker);
  }

  approveEditRow(worker: Worker) {
    const element_id = document.getElementById("id2"+worker.id) as HTMLInputElement;
    const element_full_name = document.getElementById("full_name2"+worker.id) as HTMLInputElement;
    const element_age = document.getElementById("age2"+worker.id) as HTMLInputElement;
    const element_sex = document.getElementById("sex2"+worker.id) as HTMLInputElement;
    const element_position_id = document.getElementById("position_id2"+worker.id) as HTMLInputElement;

    const new_full_name = element_full_name.value;
    const new_age = parseInt(element_age.value);
    const new_sex = element_sex.value;
    const new_position_id = parseInt(element_position_id.value);

    //full_name (6-100 chrs), age (18-120), sex ('female','male')
    if (new_full_name.length < 6 || new_full_name.length > 100) {
      alert("Ім'я повинно містити від 6 до 100 символів");
      return;
    }

    if (new_age < 18 || new_age > 120) {
      alert("Вік повинен бути від 18 до 120 років");
      return;
    }
    
    if (worker.full_name==new_full_name&&worker.age==new_age&&worker.sex==new_sex&&worker.position_id==new_position_id){
      alert("Дані не змінилися");
      return;
    }

    if (this.position_ids.indexOf(new_position_id) == -1) {
      alert("Такої посади не існує");
      return;
    }

    const new_data = {
      "full_name": new_full_name,
      "age": new_age,
      "sex": new_sex,
      "position_id": new_position_id
    };

    this.editData(worker, new_data);
    this.closeEditRow(worker);
  }

  closeEditRow(worker: Worker) {
    const element_editrow = document.getElementById("edit2"+worker.id);
    const element_editbuttons = document.getElementById("edit-buttons2"+worker.id);

    if (element_editrow && element_editbuttons) {
      element_editrow.style.display = "none";
      element_editbuttons.style.display = "";
    }

    const element_id = document.getElementById("id2"+worker.id) as HTMLInputElement;
    const element_full_name = document.getElementById("full_name2"+worker.id) as HTMLInputElement;
    const element_age = document.getElementById("age2"+worker.id) as HTMLInputElement;
    const element_sex = document.getElementById("sex2"+worker.id) as HTMLInputElement;
    const element_position_id = document.getElementById("position_id2"+worker.id) as HTMLInputElement;

    element_id.value = "";
    element_full_name.value = "";
    element_age.value = "";
    element_sex.value = "male";
    element_position_id.value = "";

    for (let i = 0; i < this.workers.length; i++) {
      if (this.workers[i].id != worker.id) {
        const element_editbuttons = document.getElementById("edit-buttons2"+this.workers[i].id);
        if (element_editbuttons){
          element_editbuttons.style.display = "";
        }
      }
    }
  }
}
