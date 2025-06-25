import { Injectable, signal } from '@angular/core';
import { of } from 'rxjs';
import { delay } from 'rxjs/operators';
import { ScrapingJob } from '../models/workflow.model';

// --- MOCK DATA ---
const MOCK_JOB: ScrapingJob = {
  id: 'job-123',
  sourceUrl: 'https://example-forum.com/thread-1',
  createdAt: new Date('2025-06-24T10:00:00Z'),
  overallStatus: 'Completed',
  workflowSteps: [
    {
      stepName: 'Web Scraping',
      status: 'Completed',
      timestamp: new Date('2025-06-24T10:01:00Z'),
      data: {
        rawHtml: '<html><body><div><p>This is some raw text with <b>bullying content</b>.</p></div></body></html>'
      }
    },
    {
      stepName: 'Data Cleaning',
      status: 'Completed',
      timestamp: new Date('2025-06-24T10:02:00Z'),
      data: {
        cleanedText: 'This is some raw text with bullying content.'
      },
      history: {
        input: '<html><body><div><p>This is some raw text with <b>bullying content</b>.</p></div></body></html>',
        output: 'This is some raw text with bullying content.'
      }
    },
    {
      stepName: 'Lemmatization',
      status: 'Completed',
      timestamp: new Date('2025-06-24T10:03:00Z'),
      data: {
        lemmatizedText: 'This be some raw text with bully content.'
      },
      history: {
        input: 'This is some raw text with bullying content.',
        output: 'This be some raw text with bully content.'
      }
    },
    {
      stepName: 'Data Enrichment',
      status: 'Completed',
      timestamp: new Date('2025-06-24T10:05:00Z'),
      data: {
        mongoDocument: {
          _id: 'mongo_doc_id_456',
          source: 'https://example-forum.com/thread-1',
          lemmatizedText: 'This be some raw text with bully content.',
          ontologyMatches: [
            { term: 'bully', category: 'AggressiveBehavior', score: 0.95 }
          ]
        }
      }
    }
  ]
};

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  // Using a signal for reactive state management
  public job = signal<ScrapingJob | null>(null);
  public isLoading = signal<boolean>(false);

  constructor() {}

  // In a real app, jobId would be passed and used in an HTTP GET request
  fetchJobData(jobId: string) {
    this.isLoading.set(true);
    this.job.set(null);

    // Simulate an API call
    of(MOCK_JOB).pipe(delay(1000)).subscribe(data => {
      this.job.set(data);
      this.isLoading.set(false);
    });
  }
}
