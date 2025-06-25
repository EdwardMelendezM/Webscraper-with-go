import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatCardModule } from '@angular/material/card';
import { ApiService } from '../../core/services/api.service';
import {StepCardComponent} from '../step-card/step-card';

@Component({
  selector: 'app-workflow-viewer',
  standalone: true,
  imports: [CommonModule, MatProgressSpinnerModule, MatCardModule, StepCardComponent],
  templateUrl: './workflow-viewer.html',
  styleUrls: ['./workflow-viewer.scss']
})
export class WorkflowViewerComponent implements OnInit {
  apiService = inject(ApiService);

  ngOnInit(): void {
    // We trigger the fetch for a specific job ID here.
    // In a real app, this ID would come from the route (e.g., /jobs/job-123)
    this.apiService.fetchJobData('job-123');
  }
}
