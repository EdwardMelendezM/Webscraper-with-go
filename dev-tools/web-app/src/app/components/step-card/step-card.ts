import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { WorkflowStep } from '../../core/models/workflow.model';

@Component({
  selector: 'app-step-card',
  standalone: true,
  imports: [CommonModule, MatCardModule, MatIconModule, MatDividerModule],
  templateUrl: './step-card.html',
  styleUrls: ['./step-card.scss']
})
export class StepCardComponent {
  @Input({ required: true }) step!: WorkflowStep;

  // Helper to format JSON data for display
  get formattedData(): string {
    return JSON.stringify(this.step.data, null, 2);
  }
}
