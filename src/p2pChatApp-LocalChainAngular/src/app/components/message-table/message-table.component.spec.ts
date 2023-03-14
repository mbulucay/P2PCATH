import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MessageTableComponent } from './message-table.component';

describe('MessageTableComponent', () => {
  let component: MessageTableComponent;
  let fixture: ComponentFixture<MessageTableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MessageTableComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MessageTableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
