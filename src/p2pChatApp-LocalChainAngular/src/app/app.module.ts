import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BlockchainViewerComponent } from './pages/blockchain-viewer/blockchain-viewer.component';
import { BlockViewComponent } from './components/block-view/block-view.component';
import { BlockchainService } from './services/blockchain.service';
import { MessageTableComponent } from './components/message-table/message-table.component';
import { SettingsComponent } from './pages/settings/settings.component';
import { SendMessageComponent } from './pages/send-message/send-message.component';
import { CreateMessageComponent } from './pages/create-message/create-message.component';
import { PendingMessagesComponent } from './pages/pending-messages/pending-messages.component'

@NgModule({
  declarations: [
    AppComponent,
    BlockchainViewerComponent,
    BlockViewComponent,
    MessageTableComponent,
    SettingsComponent,
    SendMessageComponent,
    CreateMessageComponent,
    PendingMessagesComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
  ],
  exports: [
    BlockchainViewerComponent
  ],
  providers: [BlockchainService],
  bootstrap: [AppComponent]
})
export class AppModule { }
