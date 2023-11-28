import { Component, Input} from '@angular/core';
import { HttpClient } from '@angular/common/http';

export interface Client {
  id: number;
  full_name: string;
  age: number;
  sex: string;
}

@Component({
  selector: 'app-clients-table',
  templateUrl: './clients-table.component.html',
  styleUrls: ['./clients-table.component.css']
})
export class ClientsTableComponent {
  clients: Client[] = [];

  link = 'http://localhost:8080/clients';

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.readData();
  }

  readData() {
    this.http.get<Client[]>(this.link).subscribe(data => {
      this.clients = data;
      //sort by id
      this.clients.sort((a, b) => (a.id > b.id) ? 1 : -1);
    });
  }

  editData(client_old: Client, new_data: any) {
    this.http.put<Client>(this.link + '/' + client_old.id, new_data).subscribe(data => {
      this.readData();
    }, error => {
      alert("Щось пішло не так при редагуванні даних");
    });
  }

  deleteData(client: Client) {
    this.http.delete<Client>(this.link + '/' + client.id).subscribe(data => {
      this.readData();
    }, error => {
      alert("Щось пішло не так при видаленні");
    });
  }

  addData(data: any) {
    this.http.post<Client>(this.link, data).subscribe(data => {
      this.readData();
    }, error => {
        alert("Щось пішло не так при додаванні");
      });
  }

  addRowPlaceholder() {
    const element_placeholder = document.getElementById("addRowTable");
    const element_button = document.getElementById("addRowButton");

    if (element_placeholder && element_button) {
      element_placeholder.style.display = "";
      element_button.style.display = "none";
    }
  }

  addRow() {
    const element_full_name = document.getElementById("addRowFullName") as HTMLInputElement;
    const element_age = document.getElementById("addRowAge") as HTMLInputElement;
    const element_sex = document.getElementById("addRowSex") as HTMLSelectElement;

    const new_name = element_full_name.value;
    const new_age = parseInt(element_age.value);
    const new_sex = element_sex.value;

    //full_name (6-100 chrs), age (18-120), sex ('female','male')
    if (new_name.length < 6 || new_name.length > 100) {
      alert("Ім'я повинно містити від 6 до 100 символів");
      return;
    }

    if (new_age < 18 || new_age > 120) {
      alert("Вік повинен бути від 18 до 120 років");
      return;
    }

    const new_client = {
      "full_name": new_name, 
      "age": new_age, 
      "sex": new_sex};
    this.addData(new_client);

    element_full_name.value = "";
    element_age.value = "";
    element_sex.value = "male";
  }

  closeAddRowPlaceholder() {
    const element_placeholder = document.getElementById("addRowTable");
    const element_button = document.getElementById("addRowButton");

    if (element_placeholder && element_button) {
      element_placeholder.style.display = "none";
      element_button.style.display = "";
    }

    const element_full_name = document.getElementById("addRowFullName") as HTMLInputElement;
    const element_age = document.getElementById("addRowAge") as HTMLInputElement;
    const element_sex = document.getElementById("addRowSex") as HTMLSelectElement;

    element_full_name.value = "";
    element_age.value = "";
    element_sex.value = "male";
  }

  addEditRow(client: Client) {
    const element_editrow = document.getElementById("edit"+client.id);
    const element_editbuttons = document.getElementById("edit-buttons"+client.id);

    if (element_editrow && element_editbuttons) {
      element_editrow.style.display = "";
      element_editbuttons.style.display = "none";
    }

    const element_id = document.getElementById("id"+client.id) as HTMLInputElement;
    const element_full_name = document.getElementById("full_name"+client.id) as HTMLInputElement;
    const element_age = document.getElementById("age"+client.id) as HTMLInputElement;
    const element_sex = document.getElementById("sex"+client.id) as HTMLInputElement;

    element_id.value = client.id.toString();
    element_full_name.value = client.full_name;
    element_age.value = client.age.toString();
    element_sex.value = client.sex;

    for (let i = 0; i < this.clients.length; i++) {
      if (this.clients[i].id != client.id) {
        const element_editbuttons = document.getElementById("edit-buttons"+this.clients[i].id);
        if (element_editbuttons){
          element_editbuttons.style.display = "none";
        }
      }
    }
  }

  deleteRow(client: Client) {
    this.deleteData(client);
  }

  approveEditRow(client: Client) {
    console.log("approveEditRow");
    const element_id = document.getElementById("id"+client.id) as HTMLInputElement;
    const element_full_name = document.getElementById("full_name"+client.id) as HTMLInputElement;
    const element_age = document.getElementById("age"+client.id) as HTMLInputElement;
    const element_sex = document.getElementById("sex"+client.id) as HTMLInputElement;

    const new_full_name = element_full_name.value;
    const new_age = parseInt(element_age.value);
    const new_sex = element_sex.value;

    //full_name (6-100 chrs), age (18-120), sex ('female','male')
    if (new_full_name.length < 6 || new_full_name.length > 100) {
      alert("Ім'я повинно містити від 6 до 100 символів");
      return;
    }

    if (new_age < 18 || new_age > 120) {
      alert("Вік повинен бути від 18 до 120 років");
      return;
    }
    
    if (client.full_name==new_full_name&&client.age==new_age&&client.sex==new_sex){
      alert("Дані не змінилися");
      return;
    }

    const new_data = {
      "full_name": new_full_name,
      "age": new_age,
      "sex": new_sex
    };

    this.editData(client, new_data);
    this.closeEditRow(client);
  }

  closeEditRow(client: Client) {
    const element_editrow = document.getElementById("edit"+client.id);
    const element_editbuttons = document.getElementById("edit-buttons"+client.id);

    if (element_editrow && element_editbuttons) {
      element_editrow.style.display = "none";
      element_editbuttons.style.display = "";
    }

    const element_id = document.getElementById("id"+client.id) as HTMLInputElement;
    const element_full_name = document.getElementById("full_name"+client.id) as HTMLInputElement;
    const element_age = document.getElementById("age"+client.id) as HTMLInputElement;
    const element_sex = document.getElementById("sex"+client.id) as HTMLInputElement;

    element_id.value = "";
    element_full_name.value = "";
    element_age.value = "";
    element_sex.value = "male";

    for (let i = 0; i < this.clients.length; i++) {
      if (this.clients[i].id != client.id) {
        const element_editbuttons = document.getElementById("edit-buttons"+this.clients[i].id);
        if (element_editbuttons){
          element_editbuttons.style.display = "";
        }
      }
    }
  }
}
