import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {WorkflowViewerComponent} from './components/workflow-viewer/workflow-viewer';
import {MatToolbar} from '@angular/material/toolbar';

@Component({
  selector: 'app-root',
  imports: [WorkflowViewerComponent, MatToolbar],
  templateUrl: './app.html',
  styleUrl: './app.scss'
})
export class App {
  protected title = 'web-app';
}
