import { Component } from '@angular/core';
import {
  faChevronDown,
  faChevronUp,
  faFile,
  faGear,
  faHome,
  faPaste,
  faTags,
  faUserLarge,
} from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'looped-cms-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss'],
})
export class SidebarComponent {
  homeIcon = faHome;
  postsIcon = faPaste;
  expandIcon = faChevronDown;
  nonExpandIcon = faChevronUp;
  tagsIcon = faTags;
  usersIcon = faUserLarge;
  settingsIcon = faGear;
  pagesIcon = faFile;

  isPostsVisible = false;

  isPagesVisible = false;

  constructor() {
    console.log('Sidebar');
  }

  togglePostsVisible() {
    this.isPostsVisible = !this.isPostsVisible;
  }

  togglePagesVisible() {
    this.isPagesVisible = !this.isPagesVisible;
  }
}
