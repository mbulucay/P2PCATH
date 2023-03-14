import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { BlockchainViewerComponent } from './pages/blockchain-viewer/blockchain-viewer.component';
import { CreateMessageComponent } from './pages/create-message/create-message.component';
import { PendingMessagesComponent } from './pages/pending-messages/pending-messages.component';
import { SettingsComponent } from './pages/settings/settings.component'

const routes: Routes = [
  {path: '', component: BlockchainViewerComponent},
  {path: 'settings', component: SettingsComponent},
  {path: 'new/message', component: CreateMessageComponent},
  {path: 'new/message/pending', component: PendingMessagesComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
