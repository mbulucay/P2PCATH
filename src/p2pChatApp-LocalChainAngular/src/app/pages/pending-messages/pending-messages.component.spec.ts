import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PendingMessagesComponent } from './pending-messages.component';

describe('PendingMessagesComponent', () => {
  let component: PendingMessagesComponent;
  let fixture: ComponentFixture<PendingMessagesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PendingMessagesComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PendingMessagesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
