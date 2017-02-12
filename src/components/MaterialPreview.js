import React from 'react'
import Relay from 'react-relay'
import classes from './MaterialPreview.css'

class MaterialPreview extends React.Component {

  static propTypes = {
    material: React.PropTypes.object,
    router: React.PropTypes.object,
  }

  render () {
    return (
      <div className={classes.link}>
        <div className={classes.previewPage}>
          <img className={classes.previewImg} src={this.props.material.cover} alt='cover Image' />
          <div className={classes.previewName}>
            {this.props.material.name}
          </div>
          <div className={classes.previewName}>
            {this.props.material.url}
          </div>
        </div>
      </div>
    )
  }
}

export default Relay.createContainer(
  MaterialPreview,
  {
    fragments: {
      material: () => Relay.QL`
        fragment on Material {
          cover
          name
          url
        }
      `,
    },
  }
)
