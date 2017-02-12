import React from 'react'
import Relay from 'react-relay'
import MaterialPreview from '../components/MaterialPreview'
import classes from './ListPage.css'

class ListPage extends React.Component {
  static propTypes = {
    viewer: React.PropTypes.object,
  }
  render () {
    return (
      <div className={classes.root}>
        <div className={classes.title}>
          {`当前分类 ${this.props.viewer.name} , 一共有 ${this.props.viewer.current} 个，
           全部有 ${this.props.viewer.totalCount} 个。`}
        </div>
        <div className={classes.container}>
          {this.props.viewer.materials.edges.map((edges) => edges.node).map((material) =>
            <MaterialPreview key={material.id} material={material} />
          )
          }
        </div>
      </div>
    )
  }
}

export default Relay.createContainer(
  ListPage,
  {
    fragments: {
      viewer: () => Relay.QL`
        fragment on Category {
          name,
          current,
          totalCount,
          materials(first: 10) {
              edges {
                  node {
                    id,
                    ${MaterialPreview.getFragment('material')},
                  },
              },
          },
        }
      `,
    },
  },
)
